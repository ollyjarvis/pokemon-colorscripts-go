# Maintainer: Oliver Jarvis <hello at ollyjarvis dot uk>
pkgname=pokemon-colorscripts-go
pkgver=0.0.4
pkgrel=2
pkgdesc='pokemon-colorscripts rewritten in Go'
arch=('x86_64')
url="https://github.com/ollyjarvis/$pkgname"
license=('MIT')

makedepends=('go')
depends=('glibc')

source=("$pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('78c0eeb2d0217046e02263699f150456dd94dfb93765c2cb7971a8b22ce55432')

build() {
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"
  cd "$pkgname-$pkgver"
  go build -o $pkgname -ldflags "-X main.PROGRAM_DIR=/usr/share/$pkgname" . 
}

package() {
  cd "$pkgname-$pkgver"
  mkdir -p "$pkgdir/usr/share/$pkgname/colorscripts/small/regular"
  mkdir -p "$pkgdir/usr/share/$pkgname/colorscripts/small/shiny"
  mkdir -p "$pkgdir/usr/share/$pkgname/colorscripts/large/regular"
  mkdir -p "$pkgdir/usr/share/$pkgname/colorscripts/large/shiny"
  install -Dm644 colorscripts/small/regular/* "$pkgdir/usr/share/$pkgname/colorscripts/small/regular"
  install -Dm644 colorscripts/small/shiny/* "$pkgdir/usr/share/$pkgname/colorscripts/small/shiny"
  install -Dm644 colorscripts/large/regular/* "$pkgdir/usr/share/$pkgname/colorscripts/large/regular"
  install -Dm644 colorscripts/large/shiny/* "$pkgdir/usr/share/$pkgname/colorscripts/large/shiny"
  install -Dm644 pokemon.json "$pkgdir/usr/share/$pkgname/pokemon.json"
  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
  install -Dm755 $pkgname "$pkgdir"/usr/bin/$pkgname
}
