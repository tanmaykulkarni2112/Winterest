import type { JSX } from "react";

type elements = {
    text:string
};

const Subheader = ({text}:elements):JSX.Element =>{
    return(
        <div className="text-xl text-center">
            {text}
        </div>
    )
}

export default Subheader;
