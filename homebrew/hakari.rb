require 'formula'

HOMEBREW_HAKARI_VERSION = '0.1.0'
class Hakari < Formula
  homepage 'https://github.com/showwin/hakari'
  url 'https://github.com/showwin/hakari.git', tag: "v#{HOMEBREW_HAKARI_VERSION}"
  version HOMEBREW_HAKARI_VERSION
  head 'https://github.com/showwin/hakari.git', branch: 'master'

  depends_on 'go' => :build
  depends_on :hg => :build

  def install
    ENV['GOPATH'] = buildpath
    system 'go', 'get', 'gopkg.in/yaml.v2'
    system 'go', 'build', '-o', 'hakari'
    bin.install 'hakari'
  end
end
