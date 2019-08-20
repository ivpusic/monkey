let closure = fn(x) {
    return fn (y) {
        return x + y;
    };
};

puts(closure(10)(20));