#include <iostream>
#include <vector>
using namespace std;


vector<int> first_half(vector<int> a) {
  if (a.size() < 2) {
    return a;
  }

  int len =  a.size()/ 2;
  vector<int> result;
  int i;
  for (i = 0; i < len; i++ ){
    result.push_back(a[i]);
  }

  return result;
}
vector<int> second_half(vector<int> a) {
  if (a.size() <2) {
    return a;
  }

  int len =  a.size()/ 2;
  vector<int> result;
  int i;
  for (i = len; i < a.size(); i++ ){
    result.push_back(a[i]);
  }
  return a;
}

vector<int> merge(vector<int> a, vector<int> b) {
  vector<int> result;
  if (a.size() < 2 && b.size() < 2 ) {
    // Do max and min?
    if (a[0] > b[0]) {
      result.push_back(a[0]);
      result.push_back(b[0]);
      return result;
    }
    result.push_back(b[0]);
    result.push_back(a[0]);
    return result;
  }

  vector<int> first_a;
  vector<int> second_a;
  first_a = first_half(a);
  second_a = second_half(a);
  vector<int> first_b;
  vector<int> second_b;
  first_b = first_half(b);
  second_b = second_half(b);
  vector<int> first;
  vector<int> second;
  first = merge(first_a, first_b);
  second = merge(first_b, first_b);
  return merge(first, second);


}

int main() {
  vector<int> nums;
  nums = {3, 1, 5, 13, 4, 17, 9, 16};
    //int nums [8] = {3, 1, 5, 13, 4, 17, 9, 16};

  vector<int> first;
  vector<int> second;
  first = first_half(nums);
  second = second_half(nums);
  vector<int> result = merge(first, second);
  cout << result[0];

}
