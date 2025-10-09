package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	http.HandleFunc("/shopping_cart", cartHandler)
	http.Handle("/", http.FileServer(http.Dir("."))) // 提供静态文件服务

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

type CartItem struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"` // 以元为单位
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		showCart(w, r)
	case "POST":
		addToCart(w, r)
	case "DELETE":
		removeFromCart(w, r)
	}
}

func showCart(w http.ResponseWriter, r *http.Request) {
	cartCookie, err := r.Cookie("shopping_cart")
	if err != nil {
		fmt.Fprintf(w, `{"items": [], "total": 0}`)
		return
	}

	var items []CartItem
	cookieValue, _ := url.QueryUnescape(cartCookie.Value)
	json.Unmarshal([]byte(cookieValue), &items)

	total := 0
	for _, item := range items {
		total += item.Quantity * item.Price
	}

	response := map[string]interface{}{
		"items": items,
		"total": total,
	}

	jsonData, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func addToCart(w http.ResponseWriter, r *http.Request) {
	var newItem CartItem
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 读取现有购物车
	var items []CartItem
	if cookie, err := r.Cookie("shopping_cart"); err == nil {
		cookieValue, _ := url.QueryUnescape(cookie.Value)
		json.Unmarshal([]byte(cookieValue), &items)
	}

	// 检查是否已存在，如果存在则更新数量
	found := false
	for i, item := range items {
		if item.ProductID == newItem.ProductID {
			items[i].Quantity += newItem.Quantity
			found = true
			break
		}
	}

	if !found {
		items = append(items, newItem)
	}

	// 保存到 Cookie
	itemsJSON, _ := json.Marshal(items)
	encodedValue := url.QueryEscape(string(itemsJSON))
	cookie := &http.Cookie{
		Name:     "shopping_cart",
		Value:    encodedValue,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour), // 7天有效期
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func removeFromCart(w http.ResponseWriter, r *http.Request) {
	// 解析 product_id 参数（支持 JSON 或 URL 查询参数）
	var productID string
	if r.Header.Get("Content-Type") == "application/json" {
		var req struct {
			ProductID string `json:"product_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		productID = req.ProductID
	} else {
		productID = r.URL.Query().Get("product_id")
	}

	if productID == "" {
		http.Error(w, "Missing product_id", http.StatusBadRequest)
		return
	}

	// 读取现有购物车
	var items []CartItem
	if cookie, err := r.Cookie("shopping_cart"); err == nil {
		cookieValue, _ := url.QueryUnescape(cookie.Value)
		json.Unmarshal([]byte(cookieValue), &items)
	}

	// 移除指定商品
	newItems := make([]CartItem, 0, len(items))
	for _, item := range items {
		if item.ProductID != productID {
			newItems = append(newItems, item)
		}
	}

	// 保存到 Cookie
	itemsJSON, _ := json.Marshal(newItems)
	encodedValue := url.QueryEscape(string(itemsJSON))
	cookie := &http.Cookie{
		Name:     "shopping_cart",
		Value:    encodedValue,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
