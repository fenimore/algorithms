#lang racket
(newline)
;; I like newline


;; merge sort
(define (split x) ;; split is a function of one arg
  "Split will split a list into two lists"
  ;; TODO: Doesn't work with odd lists
  (define mididx
    (if
     (odd? (length x))
     ("Doesn't work")  ;;((/ (- (length x) 1) 2))
     (/ (length x) 2))
    )
  ;; Get first and second half of list
  (list (take x  mididx) (drop x mididx))
  )



(define (merge a b c)
  "Merge merges two already sorted list"
  (append c (list
  (cond
   ((< (car a) (car b)) (car a))
   ((< (car b) (car a)) (car b));;(append c (list (car a)))
   )
  ))
  (define x "a")
  (define y "b")
  (c)
  )


;
(display (split '(1 5 4 2)))
;;(display (merge (list 1 2 3) (list 6 5 4) '()))
;; I can do (list n n n) or literal: `(n n n)

;;(display (merge `(1 2 3) `(6 5 4)
(merge (list 1 2 3) (list 6 5 4) '())



(newline)
