var express = "0";
var ans = "0";
function calculate(val) {
    var result = document.getElementById("result");
    switch(val) {
    case "c":
        express = "0";
        ans = "0";
        result.innerHTML = express;
        break;
    case "d":
        express = (express != "0") ? express : "";
        express = express.substr(0, express.length - 1);
        express = (express != "") ? express : "0";
        result.innerHTML = express;
        break;
    case "=":
        try {
            ans = eval(express);
            if(ans == Infinity)
                result.innerHTML = "不能除以0";
            else {
                if(ans.toString().length > 9)
                ans = ans.toPrecision(9);
                result.innerHTML = "=" + " " + ans;   
            }
        }
        catch (err){
            result.innerHTML = "错误表达式";
            alert("错误表达式");  
        }
        express = "0"; 
        break;
    case "+":
    case "-":
    case "*":
    case "/":
    case "%":
    case ".":
        if(express == "0" && val != ".") {
            express = ans;
            express += val;
            result.innerHTML = express;
        }
        else if(express[express.length - 1] != '+' && express[express.length - 1] != '-' && express[express.length - 1] != '*' && express[express.length - 1] != '/' && express[express.length - 1] != '.' && express[express.length - 1] != '%') {
            express += val;
            result.innerHTML = express;
        }
        try {
            ans = eval(express);
        }
        catch (err){
            ans = ans;
        }
        break;
    case "n":
        try {
            ans = -1 * ans; 
            express = ans;
            result.innerHTML = ans;
        }
        catch (err){
            result.innerHTML = "错误表达式";
            alert("错误表达式");  
        }
        break;
    default:
        if(express != "0"){
            express += val;
        }
        else
            express = val; 
        result.innerHTML = express;
        try {
            ans = eval(express);
        }
        catch (err){
            ans = ans;
        }
    }    
    if(result.innerHTML.length <= 13) {
        result.style.fontSize = 60+'px';
    }
    else if(result.innerHTML.length <= 14){
        result.style.fontSize = 54 +'px';
    }
    else if(result.innerHTML.length <= 15){
        result.style.fontSize = 50 +'px';
    }
    else {
        result.innerHTML = "超出范围";
    }
}