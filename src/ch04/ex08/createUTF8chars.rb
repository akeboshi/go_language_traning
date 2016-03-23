#! /usr/bin/env ruby
# 適当なutf8文字列を、適当な数出力します
codes = (0x00 ... 0x1ffff).to_a
chars = ''
1000.times do
  begin
    chars += codes[rand(codes.length)].chr("UTF-8")
  rescue
  end
end
print chars
