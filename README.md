<div align="center">
<img src="assets/cvforge.png" alt="CVForge" width="200">
<h1>CVForge</h1>

<p>
A lightweight Computer Vision CLI toolkit written in Go.
</p>

<a href="https://go.dev/">
  <img src="https://img.shields.io/badge/go-1.22+-blue.svg?style=for-the-badge&logo=go" alt="Go Version">
</a>

<a href="LICENSE">
  <img src="https://img.shields.io/badge/license-MIT-green.svg?style=for-the-badge" alt="License">
</a>

</div>


---

**CVForge** 是一個以 Go 開發的輕量級 Computer Vision CLI 工具。

透過模組化 Filter Pipeline 架構，讓使用者可以快速組合影像處理流程，
從基礎影像轉換到進階電腦視覺演算法皆可擴充。

---

## 核心特色

* **Modular Filter Pipeline**

  每個影像處理功能皆為獨立 Filter，可自由組合：

```text
Input Image
     |
     v
+-----------+
| Grayscale |
+-----------+
     |
     v
+-----------+
|   Blur    |
+-----------+
     |
     v
+-----------+
|  Resize   |
+-----------+
     |
     v
Output Image
```
## 已支援功能
| Feature   | Command     | Description        |
| --------- | ----------- | ------------------ |
| Grayscale | `grayscale` | RGB 轉灰階            |
| Blur      | `blur`      | Gaussian-like Blur |
| Resize    | `resize`    | 調整影像尺寸             |
| Rotate    | `rotate`    | 旋轉影像               |
| Crop      | `crop`      | 裁切影像               |

## 快速開始
### 安裝
#### Clone repository:
```aiignore
git clone https://github.com/kylin419/CVforge.git
cd CVforge
```
#### Build:
```aiignore
go build .
```
### 使用方式
#### GrayScale
```aiignore
cvforge grayscale input.jpg output.jpg
```
#### Resize
```aiignore
cvforge resize \
input.jpg \
output.jpg \
--w WIDTH \
--h  HEIGHT
```
#### Rotate
```aiignore
cvforge rotate \
input.jpg \
output.jpg \
--angle ANGLE
```

#### Crop
```aiignore
cvforge crop \
input.jpg \
output.jpg \
--x X \
--y Y \
--w W \
--h H
```

## License
### MIT License

## Author
### Kylin
