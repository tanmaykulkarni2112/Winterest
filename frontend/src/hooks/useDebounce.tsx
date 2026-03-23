import { useEffect, useState } from "react";

const useDebounce=(val:string,delay:number):string=>{
    const [debouncedVal,setDebouncedVal]=useState<string>("");

    useEffect(()=>{
        const id = setTimeout(()=>{
            setDebouncedVal(val);
        },delay)
        
        return(()=>{
            clearTimeout(id);
        })
    },[val,delay])

    return debouncedVal;
}

export default useDebounce;