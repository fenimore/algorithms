#lang racket
(newline)


;; merge sort

;; Split
(define (split x) ;; split is a function of one arg
  "Split will split a list into two lists"
  ;; TODO: Doesn't work with odd lists
  (define mididx
    (if
     (odd? (length x))
     ("Doesn't work on odd lists")  ;;((/ (- (length x) 1) 2))
     (/ (length x) 2))
    )
  ;; Get first and second half of list
  (list (take x  mididx) (drop x mididx))
  )



(define (merge a b c)
   "Merge merges two already sorted list"
   ;;(define first-a (car a)) ;; first a
   ;;(define first-b (car b)) ;; first b
   ;; Add the lowest to c list
   (let* (
          [z
           (append c
                   (cond
                    ((> (car a) (car b))(car b))
                    ((< (car a) (car b))(car a))
                    )
                   )]
          ;; Pop the item off such a list
          ;; x is the popped list
          [x
           (cond
            ((> (car a) (car b))(drop b 1))
            ((< (car a) (car b))(drop a 1))
            )]
          ;; The list not popped
          [y
           (cond
            ((> (car a) (car b)) a)
            ((< (car a) (car b)) b)
            )]
          )
     ;; Call merge if the lists are not 0
     ;; pass in x for the popped list
     ;; Emptiness is true if the poped list is empy
     ;;(display '(x y z))
     (display (list? z))
     (display x)
     (display y)
     (newline)
     (newline)
     (if
      (null? x)
      (append z y)
      (merge x y (list z))
      )
     )
   )




;;(display (split '(1 5 4 2)))
;;(merge (list 1 2 3) (list 6 5 4) '())
(merge '(1 5 7 11) '(4 6 8 10) '())



(newline)
