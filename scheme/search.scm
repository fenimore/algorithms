#lang racket

(define sorted-list '(2 4 5 6 8 14 15 17))


(define (binary-search lst target high low)
  "Binary Search for index of target number"
  ;; Always takes length of list for high and
  ;; 0 for low
  ;; TODO: If number outside of range, fail
  (define mididx
    (if
     (odd? (+ low high))
     (/ (- (+ low high) 1) 2)
     (/ (+ low high) 2)))

  (cond
   ((> target (list-ref lst (- high 1))) "Not Found")
   ((= (list-ref lst mididx) target) mididx)
   ((< target (list-ref lst mididx)) (binary-search lst target mididx low))
   ((> target (list-ref lst mididx)) (binary-search lst target high mididx))
   )
  )

(display "Search For 5 ")
(binary-search sorted-list 5 (length sorted-list) 0)

(display "Search For 15 ")
(binary-search sorted-list 15 (length sorted-list) 0)

(display "Search For 33 ")
(binary-search sorted-list 33 (length sorted-list) 0)

(display "Search For 3 ")
(binary-search sorted-list 3 (length sorted-list) 0)
