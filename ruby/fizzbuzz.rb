#!env ruby

def fb(a, b, x)
  if x % a == 0 && x % b == 0
    'FB'
  elsif x % a == 0
    'F'
  elsif x % b == 0
    'B'
  else
    x.to_s
  end
end

IO.foreach(ARGV.first) do |l|
  (a, b, n) = l.split.map(&:to_i)
  puts((1..n).map { |x| fb(a, b, x) }.join(' '))
end
