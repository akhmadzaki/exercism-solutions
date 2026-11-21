#include "reverse_string.h"
#include<algorithm>

namespace reverse_string {

    std::string reverse_string(std::string s) {
        std::string rev="";
        for (std::string::reverse_iterator rit=s.rbegin(); rit!=s.rend(); ++rit)
               rev.push_back(*rit);
        return rev;
    }

}  // namespace reverse_string
