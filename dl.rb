require 'net/http'
require 'json'
require 'clipboard'


def get_gist_url(name)
    response = Net::HTTP.get_response(URI('https://git.io/' + name))
    gist_url = response['location']
    return gist_url
end


def get_gist_code(gist_url)
    gist = Net::HTTP.get_response(URI(gist_url))
    code = gist.body
    return code
end


def download_code(name)
    gist_url = get_gist_url(name)
    filename =  gist_url.split('/').last

    File.open(filename, 'w') { |file| file.write(code) }
    return filename
end

## Uploading Code

# name = 'dl.rb'
# code = File.open(name, 'rb').read
# uri = URI('https://api.github.com/gists')
# payload = {
#   'files' => {
#     name => {
#       'content' => code
#     }
#   }
# }
# req = Net::HTTP::Post.new(uri.path)
# req.body = payload.to_json
# res = Net::HTTP.start(uri.hostname, uri.port, :use_ssl => true) do |http|
#   http.request(req)
# end
# puts res.body.inspect

# uri = URI('https://git.io')
# res = Net::HTTP.post_form(uri, 'url' => 'https://github.com/thisissocold', 'code' => name)
# puts res.code
# puts res['location']
