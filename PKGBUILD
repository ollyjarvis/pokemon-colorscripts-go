# Maintainer: Oliver Jarvis <hello@ollyjarvis.uk>
pkgname=pokemon-colorscripts-go
pkgver=0.0.2
pkgrel=1
pkgdesc='pokemon-colorscripts rewritten in Go'
arch=('x86_64')
url="https://github.com/ollyjarvis/$pkgname"
license=('MIT')

makedepends=('go')
depends=('glibc')

source=("$pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('15012306fff0a20c382e56be4edf071739d876b8416e255c8d575f3d6e9f4380')

build() {
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"
  cd "$pkgname-$pkgver"
  go build -o $pkgname .
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