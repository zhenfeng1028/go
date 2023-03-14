# sha
安全散列算法SHA（Secure Hash Algorithm）是美国国家安全局 （NSA） 设计，美国国家标准与技术研究院（NIST） 发布的一系列密码散列函数，包括 SHA-1、SHA-224、SHA-256、SHA-384 和 SHA-512 等变体。主要适用于数字签名标准（DigitalSignature Standard DSS）里面定义的数字签名算法（Digital Signature Algorithm DSA）。SHA-1已经不是那边安全了，google和微软都已经弃用这个加密算法。为此，我们使用热门的比特币使用过的算法SHA-256作为实例。其它SHA算法，也可以按照这种模式进行使用。

哈希值用作表示大量数据的固定大小的唯一值。数据的少量更改会在哈希值中产生不可预知的大量更改。 SHA256 算法的哈希值大小为 256 位。