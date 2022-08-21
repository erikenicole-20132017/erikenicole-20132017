def memoize(func)
  cache = {}
  lambda { |*args| cache[args] ||= func.call(*args) }
end
