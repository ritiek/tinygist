require "tinygist/version"
require 'net/http'
require 'json'

module Tinygist
  def get_gist_url(name)
    response = Net::HTTP.get_response(URI('https://git.io/' + name))
    gist_url = response['location']
    return gist_url
  end

  def raw_code(gist_url)
    gist = Net::HTTP.get_response(URI(gist_url))
    code = gist.body
    return code
  end

  def create_gist(name, code)
    uri = URI('https://api.github.com/gists')
    payload = {
      'files' => {
        name => {
          'content' => code
        }
      }
    }

    req = Net::HTTP::Post.new(uri.path)
    req.body = payload.to_json
    res = Net::HTTP.start(uri.hostname, uri.port, :use_ssl => true) do |http|
      http.request(req)
    end

    raw_url = JSON.parse(res.body)['files'][name]['raw_url']
    return raw_url
  end

  def shorten_url(github_url, code)
    uri = URI('https://git.io')
    response = Net::HTTP.post_form(uri, 'url' => github_url, 'code' => code)
    return response
  end
end
