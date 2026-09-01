require 'uri'
require 'net/http'

url = URI("https://api-sandbox.capitalos.com/accounts/act_01KV90QQAGQ8VPWPSNG7B7KNQK/cards/")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Post.new(url)
request["accept"] = 'application/json'
request["content-type"] = 'application/json'
request["Authorization"] = 'cos_sk_sbx_2af654c2-06f3-47e3-9ac3-47b6f7b19dc7'
request["x-capitalos-on-behalf-of-user-id"] = 'act_user_01KVB2NXAA55W1SFKM85Q1351P'
request.body = "{\"cardholder\":{\"firstName\":\"rafael\",\"lastName\":\"macedo\"},\"accountUserId\":\"act_user_01KVTCFJ31NHJA0X95D05QDWQR\"}"

response = http.request(request)
puts response.read_body
