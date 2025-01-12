class Lca < Formula
  desc "LCA: An alternative to ls that computes file hashes, provides color-coded, user-friendly output, and supports configurable recursion depth."
  homepage "https://github.com/dacalin/lca"
  url "https://github.com/dacalin/lca/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "your_tarball_sha256_here"
  license "MIT"

  def install
    system "go", "build", "-o", "lca"
    bin.install "lca"
  end

  test do
    system "#{bin}/lca", "-h"
  end
end
