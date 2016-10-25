;; guile
;; This fibonacci isn't implemented correctly.

(define (fib y)
  (if
   (< y 2)
   1
   ("what?")
   ;;(fib (* x (- x 1)))
   )
  )


(display (fib 8))
