你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# *Sorting-Algorithms*


## 1. Bubble-Sort


### Do you know what the Bubble Sort algorithm is? 


The Bubble Sort algorithm is a simple **order** algorithm. It compares the next number, and if the number is smaller,
it changes places until the next number is larger than itself. Scroll through the numbers until they are all in order.


#### Example 


![Bubble Sort](https://user-images.githubusercontent.com/47563193/65832810-4497fe80-e29f-11e9-9fb6-138e2876fda7.png)  


The idea is to cycle the vector several times, and with each pass float to the **top** the largest element of the sequence. This movement resembles how bubbles in a water tank seek their own level, and hence the name of the algorithm.  


Class: *Classification algorithm*  
Data structure: *Matrix, Selected lists*  
Worst case complexity: *O(n²)*  
Medium case complexity: *O(n²)*  
Complexity best case: *O(n)*  
Complexity of spaces worst *case: O(n)*  


##


##  2. Insertion-Sort


### Do you know what the Insertion Sort algorithm is? Let's go to example! 


#### Playing cards


Imagine that you are playing deck. You have the cards in hand, and they are in order.
You get exactly one new card. You must place it in the correct position so that the cards in your hand remain in order.
You get another new card, it may be smaller than some of the cards you are already holding, and then you examine them, comparing the new 
card with all the cards in your hand until you find the position to put it. You insert the new card in the correct position, and again 
your hand is made up of fully ordered cards. Then you receive another card and repeat the same procedure.
Then another card, and another, and so on, until you get no more letters.  


![insertion-sort exemplo baralho](https://user-images.githubusercontent.com/47563193/65826694-eac41400-e25f-11e9-9b6b-e511599c4cda.jpg)  


The Insertion Sort starts working with the second value of the vector and then throws it to the left (beginning of the vector).
It travels the entire vector only once, but to make the movement described (play to the beginning) it uses another internal loop.
The scheme would look like this ...  


![insertionsort Exemplo](https://user-images.githubusercontent.com/47563193/65826667-9d47a700-e25f-11e9-9921-967092908b99.png)  


Class:  *Classification algorithm*  
Data structure: *Matrix, Selected lists*  
Worst case complexity: *O(n²)*  
Medium case complexity: *O(n²)*  
Complexity best case: *O(n)*  
Complexity of spaces worst case: *O(n) total, O(1) help*  
Stable: *Stability*  


###### Algorithm developed for studies only.
