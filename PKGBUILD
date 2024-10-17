pkgname=pokemon-colorscripts-go
pkgver=0.0.1
pkgrel=1
pkgdesc='pokemon-colorscripts rewritten in Go'
arch=('x86_64')
url="https://github.com/ollyjarvis/$pkgname"
license=('MIT')
makedepends=('go')
source=("$pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('708f61a49886832b3e7f7ba74dae24b27c17a7150de4e7e67cd0e43ed093a97a')


build(){
  cd "$pkgname-$pkgver"
  go build -o $pkgname .
}

package() {
  cd "$pkgname-$pkgver"
  install -Dm755 $pkgname "$pkgdir"/usr/bin/$pkgname
}
