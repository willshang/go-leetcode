
# leetcode194_转置文件
awk '{
    for (i=1;i<=NF;i++){
        if (NR==1){
            res[i]=$i
        }
        else{
            res[i]=res[i]" "$i
        }
    }
}END{
    for(i=1;i<=NF;i++){
        print res[i]
    }
}' file.txt