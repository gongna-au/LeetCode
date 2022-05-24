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
        
}
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

int main(){

   auto t = chrono::system_clock::to_time_t(std::chrono::system_clock::now());
	std::cout  << std::put_time(std::localtime(&t), "%F %T") << std::endl;
 
	//转为字符串
	std::stringstream ss;
	ss << std::put_time(std::localtime(&t), "%F %T");
	std::string str = ss.str();
 
	return 0;


}