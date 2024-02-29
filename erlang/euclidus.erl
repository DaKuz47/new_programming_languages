-module(euclidus).
-export([main/1]).

main([A, B]) ->
  R = gcd(list_to_integer(atom_to_list(A)), list_to_integer(atom_to_list(B))),
  io:format("GCD is ~w", [R]).

gcd(A, B) when B == 0 -> A;

gcd(A, B) ->
  MAX_VAL = max(A, B),
  MIN_VAL = min(A, B),
  gcd(MIN_VAL, MAX_VAL rem MIN_VAL).

