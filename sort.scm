(newline)
;; I like newline


;; merge sort
" I need to define a split and a merge "
" and then merge_sort"

(define (split x) ;; split is a function of one arg
  "Split will split a list into two lists"
  (define middle (/ (length x) 2))
  (define (before? x) (< x middle)) ;; this deosn't make sense
  ;; first half:
  (filter ;; returns a list of elements for which item returns true
   before? x)
  ;;(list 0 middle) ;; create new lists...
  )

(define (merge a b c)
  (cond
   ((< (car a) (car b)) (append c (list (car a))))
   ((> (car a) (car b)) (append c (list (car b))))
   )
  (display c)
  (if
   (not (or (null? b) (null? a)))
   (merge a b c)
   (display "ye"))
  )

(define (t a b c)
  (if
   (< (car a) (car b))
   (append c (list (car a)))
   )

;;s  (if
  ;;(display (> (car a) (car b)))
  )


;;(display (split (list 1 5 4 2)))
;;(display (merge (list 1 2 3) (list 6 5 4) '()))
(define l '())
(display (t (list 1 2 3) (list 6 5 4) (t (list 1 2 3) (list 6 5 4) l)))




(newline)
