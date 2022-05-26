#include <iostream>
#include <list>
#include <map>
#include <string>
#include <iostream>
#include <sstream>
#include <iomanip>
#include <chrono>
#include<typeinfo>


using namespace std;

bool InputIsLegal(string legalOptions,string bePressed){

    string::size_type idx;
     idx=legalOptions.find(bePressed); 
     if (idx == string::npos ){
         return false;
     }else {
         return true;
    }
        
};
class Temp{
    public:
    string name;
    Temp(string str){
        name=str;
    };
};
class test{
    public:
    list<string> values;
    test(const string a[],int n){
        for (int i=0;i<n;i++){
            values.push_back(a[i]);
        }
        
    };
    void show (){
        int i=0;
        for (list<string>::iterator iter = values.begin(); iter != values.end(); ++iter) {
            cout <<" >>" <<i <<"."<<*iter << "\n";
            i++;
        }
    };

};






int main() 
{ 
    list<int> mylist{ 1, 2, 3, 4, 5 }; 
    cout << mylist.back(); 
    return 0; 
}