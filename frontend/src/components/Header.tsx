import type { JSX } from "react";

type elements ={
    text:string
}

const Header =({text}:elements):JSX.Element=>{
 return(
    <div className="text-4xl font-bold text-center">
        {text}
    </div>
 )
}

export default Header;