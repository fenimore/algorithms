(newline)
;; I like newline


;; merge sort
" I need to define a split and a merge "

(define (split x) ;; split is a function of one arg
  "Split will split a list into two lists"
  (define middle (/ (length x) 2))
  (define (before? x) (< x middle))
  ;; first half:
  (filter ;; returns a list of elements for which item returns true
   before? x)
  ;;(list 0 middle) ;; create new lists...
  )


(display (split (list 1 5 4 2)))




(newline)
