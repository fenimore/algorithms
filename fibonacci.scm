;; guile
;; This fibonacci isn't implemented correctly.

(define (fib x)
  (fib (* x (- x 1)))
  )


(display (fib 4))
