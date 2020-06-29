# myrsa
myrsa is simple implementation of RSA, very basic 

## Algorithm
Generate public key based on two prime number p & q

1. p = 2, q = 7
2. N = p.q = 14 (modular)
3. Φ(N) = (p-1)(q-1) = 6
4. choose e ⌈ 1 < e < Φ(N)
            ⌊ coprime with N & Φ(N)
          e = 5
5. choose d; de | Φ(N) = 1
             5d | 6 = 1
             d = 5 or 14
