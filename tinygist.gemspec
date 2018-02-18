
lib = File.expand_path("../lib", __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require "tinygist/version"

Gem::Specification.new do |spec|
  spec.name          = "tinygist"
  spec.version       = Tinygist::VERSION
  spec.authors       = ["ritiek"]
  spec.email         = ["ritiekmalhotra123@gmail.com"]

  spec.summary       = "A simple tool to to share gist content from and to clipboard."
  spec.homepage      = "https://github.com/ritiek/tinygist"
  spec.license       = "MIT"

  spec.files         = `git ls-files -z`.split("\x0").reject do |f|
    f.match(%r{^(test|spec|features)/})
  end
  spec.bindir        = "exe"
  spec.executables   = spec.files.grep(%r{^exe/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]

  spec.add_dependency "clipboard", "~> 1.1"

  spec.add_development_dependency "bundler", "~> 1.16"
  spec.add_development_dependency "rake", "~> 10.0"
end
