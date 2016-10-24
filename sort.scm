#lang racket
;; merge sort
;; TODO: If duplicate numbers
;; TODO: doesn't work of zeros
;; TODO: doesn't work with odd lists

;; Split
(define (split x) ;; split is a function of one arg
  "Split will split a list into two lists"
  (define mididx
    (if
     (odd? (length x))
     (/ (- (length x) 1) 2)
     (/ (length x) 2))
    )
  ;; Get first and second half of list
  (list (take x  mididx) (drop x mididx))
  )

;; merge
(define (merge a b c)
   "Merge merges two already sorted list"
   ;;(define first-a (car a)) ;; first a
   ;;(define first-b (car b)) ;; first b
   ;; Add the lowest to c list
   (if
    ;; Merge if not list, just
    (not (list? a))
    (if
     (< a b)
     (list a b)
     (list b a)
     )
    (let*
        (
         [z
          (append c ;; append takes list
                  (list
                   (cond
                    ;;((< (length a) 2)(car b))
                    ((> (car a) (car b))(car b))
                    ((< (car a) (car b))(car a))
                    )
                   ))]
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
      (if
       (null? x)
       (append z y)
       (merge x y z)
       )
      )
    )

   )

(define (merge-sort lst)
  "Merge Sort Merges sorted lists, if list is smaller than 2
   it just returns that list"
  (if
   (< (length lst) 2)
   (car lst)
   (merge
    (merge-sort (list-ref (split lst) 0))
    (merge-sort (list-ref (split lst) 1))
    '())
   )

  )

(display "sort: ")
(display '(2 5 4 9 14 10 8 11))
(newline)
(merge-sort '(2 5 4 9 14 10 8 11))
;; TODO: doesn't work
;;(display "sort: ")
;;(display '(44 6 99 4 5 14 1 8 66 3))
;;(merge-sort '(44 6 99 4 5 14 1 8 66 3))


(newline)
