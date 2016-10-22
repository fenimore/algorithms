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
  (define first-a (car a)) ;; first a
  (define first-b (car b)) ;; first b

  ;; Add the lowest to c list
  (define z
    (append c
            ((cond
              ((> first-a first-b)(first-b))
              ((< first-a first-b)(first-a))
              )
             )
            )
    )
  ;; Pop the item off such a list
  ;; x is the popped list
  (define x
    (
     (cond
      ((> first-a first-b)(drop b 0))
      ((< first-a first-b)(drop a 0))
      )
     )
    )

  ;; The list not popped
  (define y
    (
     (cond
      ((> first-a first-b)(a))
      ((< first-a first-b)(b))
      )
     )
    )

  ;; Call merge if the lists are not 0
  ;; pass in x for the popped list
  ;; Emptiness is true if the poped list is empy
  (if
   (null? x)
   (append z y)
   (merge x y z)
   )
  )




;;(display (split '(1 5 4 2)))
;;(merge (list 1 2 3) (list 6 5 4) '())
(merge ('(1 2 3)) ('(6 5 4)) '())



(newline)
