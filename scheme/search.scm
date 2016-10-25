#lang racket

(define sorted-list '(2 4 5 6 8 14 15 17))

;; Split
(define (split x) ;; split is a function of one arg
  "Split will split a list into two lists"
  (define mididx
    (if
     (odd? (length x))
     (/ (- (length x) 1) 2)
     (/ (length x) 2)))
  ;; Get first and second half of list
  (list (take x  mididx) (drop x mididx)))

(define (binary-search lst target)
  "Binary Search for index of target number"
  ;; TODO: If number outside of range, fail
  (define mididx
    (if
     (odd? (length lst))
     (/ (- (length lst) 1) 2)
     (/ (length lst) 2)))
  (cond
   ((= (list-ref lst mididx) target) mididx)
   ((= (length lst) 1) "Item Not Found")
   ((< target (list-ref lst mididx)) (binary-search (car (split lst)) target))
   ((> target (list-ref lst mididx)) (binary-search (cdr (split lst)) target))))


(binary-search sorted-list 15)
