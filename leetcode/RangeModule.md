> A Range Module is a module that tracks ranges of numbers. Your task is to design and implement the following interfaces in an efficient manner.

- addRange(int left, int right) Adds the half-open interval [left, right), tracking every real number in that interval. Adding an interval that partially overlaps with currently tracked numbers should add any numbers in the interval [left, right) that are not already tracked.
- queryRange(int left, int right) Returns true if and only if every real number in the interval [left, right) is currently being tracked.
- removeRange(int left, int right) Stops tracking every real number currently being tracked in the interval [left, right).
# Example 1:
- addRange(10, 20): null
- removeRange(14, 16): null
- queryRange(10, 14): true (Every number in [10, 14) is being tracked)
- queryRange(13, 15): false (Numbers like 14, 14.03, 14.17 in [13, 15) are not being tracked)
- queryRange(16, 17): true (The number 16 in [16, 17) is still being tracked, despite the remove operation)
  Note:

A half open interval [left, right) denotes all real numbers left <= x < right.  
0 < left < right < 10^9 in all calls to addRange, queryRange, removeRange.  
The total number of calls to addRange in a single test case is at most 1000.  
The total number of calls to queryRange in a single test case is at most 5000.  
The total number of calls to removeRange in a single test case is at most 1000.  





# 思路 

最初对题意理解错误了，英语渣的就是这么惨。就是单纯的设置两个数组，一个是跟踪范围，一个是不跟踪范围。add和remove就是单纯的向数组里添加纪录，query就是先检测是否在不跟踪，如果在就返回false.最后检测是否在跟踪范围内，如果在就返回true,最后测试出错，才发现理解错了。



理清题义后。第一想法是申请slice，存储跟踪范围，add函数就是数组添加记录,remove函数中遍历整个slice，动态修改跟踪范围，query函数只需要检测是否符合其中一个记录即可。编码实现后，发现动态remove函数稍微好点，动态修改记录就行。add函数需要考虑太多的情况，添加后有可能有的记录范围重叠了或者重复了，需要合并，但是合并想了半天没有好的实现办法，这个想法就放弃了  



 最后想了好半天，理清了如果给定两个记录a,b，其中两者之间有六种关系，无重叠(a在b左，a在b右)，部分重叠（左重叠，右重叠）,包含（a包含b,b包含a).理清了这个，剩下就很好理解。利用map结构存储跟踪范围，key存储left,value存储right,map中所有记录不会有边界重叠的关系，都是相互独立的

 

add函数: 根据需要修改map中key内容，维护跟踪变量。两个变量min,max,一个需要删除的key数组,依次遍历map。遍历的目的就是构造一个新的key值,遍历过程中会记录需要删除的key值。最后该删除该删除，添加一个新key记录



remove函数:这个跟add函数类似，记录删除的key，构造新的key记录



query函数:遍历，只要满足其中一个即可返回true,若都不满足，则返回false 

代码见RangeModule.go