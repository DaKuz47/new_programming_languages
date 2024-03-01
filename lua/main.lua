function func(x)
    return x ^ 3 + math.exp(x)
end

function dfunc(x)
    return 3 * x ^ 2 + math.exp(x)
end

function newton(start_pos, eps)
    to_approach = func(start_pos) / dfunc(start_pos)
    next_position = start_pos - to_approach

    if math.abs(to_approach) > eps then
        return newton(next_position, eps)
    else
        return next_position
    end
end

print("Найденный корень: ")
print(newton(arg[1], 10^(-6)))
