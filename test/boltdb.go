package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db  *bolt.DB
	err error
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := &Store{
		db: db,
	}

	db.Update(func(tx *bolt.Tx) error {
		store.userBuckt, err = tx.CreateBucket([]byte("users"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return err
	})

	user := &User{Name: "lzf"}
	err = store.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("users"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%v, value=%s\n", k, v)
		}

		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("test_for_prefix"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("prefix_1"), []byte("prefix_1_value"))
		b.Put([]byte("prefix_2"), []byte("prefix_2_value"))
		b.Put([]byte("prefix_3"), []byte("prefix_3_value"))
		return err
	})

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		c := tx.Bucket([]byte("test_for_prefix")).Cursor()

		prefix := []byte("prefix_")
		k, v := c.Seek(prefix)
		fmt.Printf("key=%s, value=%s\n", k, v)
		// for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
		// 	fmt.Printf("key=%s, value=%s\n", k, v)
		// }

		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucket([]byte("Events"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return err
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Events"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("1990-01-01T00:00:00Z"), []byte("time_1"))
		b.Put([]byte("1991-01-01T00:00:00Z"), []byte("time_2"))
		b.Put([]byte("1992-01-01T00:00:00Z"), []byte("time_3"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		// Assume our events bucket exists and has RFC3339 encoded time keys.
		c := tx.Bucket([]byte("Events")).Cursor()

		// Our time range spans the 90's decade.
		min := []byte("1990-01-01T00:00:00Z")
		max := []byte("2000-01-01T00:00:00Z")

		// Iterate over the 90's.
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}

		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("users"))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%v, value=%s\n", k, v)
			return nil
		})
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucket([]byte("1000001"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return err
	})

	user = &User{Name: "ggg"}
	err = createUser(1000001, user)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/backup", BackupHandleFunc)

	log.Fatal(http.ListenAndServe(":8080", nil))

	go func() {
		// Grab the initial stats.
		prev := db.Stats()

		for {
			// Wait for 10s.
			time.Sleep(10 * time.Second)

			// Grab the current stats and diff them.
			stats := db.Stats()
			diff := stats.Sub(&prev)

			// Encode stats to JSON and print to STDERR.
			json.NewEncoder(os.Stderr).Encode(diff)

			// Save stats for the next loop.
			prev = stats
		}
	}()
	select {}
}

func BackupHandleFunc(w http.ResponseWriter, req *http.Request) {
	err := db.View(func(tx *bolt.Tx) error {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", `attachment; filename="my.db"`)
		w.Header().Set("Content-Length", strconv.Itoa(int(tx.Size())))
		_, err := tx.WriteTo(w)
		return err
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// createUser creates a new user in the given account.
func createUser(accountID int, u *User) error {
	// Start the transaction.
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Retrieve the root bucket for the account.
	// Assume this has already been created when the account was set up.
	root := tx.Bucket([]byte(strconv.FormatUint(uint64(accountID), 10)))

	// Setup the users bucket.
	bkt, err := root.CreateBucketIfNotExists([]byte("USERS"))
	if err != nil {
		return err
	}

	// Generate an ID for the new user.
	userID, err := bkt.NextSequence()
	if err != nil {
		return err
	}
	u.ID = userID

	// Marshal and save the encoded user.
	if buf, err := json.Marshal(u); err != nil {
		return err
	} else if err := bkt.Put([]byte(strconv.FormatUint(u.ID, 10)), buf); err != nil {
		return err
	}

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

type Store struct {
	db        *bolt.DB
	userBuckt *bolt.Bucket
}

// CreateUser saves u to the store. The new user ID is set on u once the data is persisted.
func (s *Store) CreateUser(u *User) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("users"))

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
		u.ID = id

		// Marshal user data into bytes.
		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(int(u.ID)), buf)
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

type User struct {
	ID   uint64
	Name string
}
