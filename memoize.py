def memoize(func):
cache = {}
def wrapper(*args):
    if args in cache:
        return cache[args]
    else:
        cache[args] = func(*args)
        return cache[args]
return wrapper
