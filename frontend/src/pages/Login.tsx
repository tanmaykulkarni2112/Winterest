import type { JSX } from "react";
import Header from "../components/Header";
import Subheader from "../components/Subheading";
import InputBoxes from "../components/InputBoxes";
import AuthBottom from "../components/AuthBottom";

type loginProps = {
    onClose: ()=>void;
}  
const LoginPage = ({onClose}:loginProps):JSX.Element =>{
    return(
        <>
            <div className="fixed inset-0 bg-black/40 backdrop-blur-sm flex items-center justify-center z-10">
                <div className="bg-white w-105 rounded-2xl shadow-2xl p-8 relative flex flex-col gap-5">
                    {/* x button */}
                    <button onClick={onClose} className="absolute right-4 top-4 text-gray-500 hover:text-black text-lg font-bold">
                        ✕
                    </button>

                    {/* heading element */}
                    <Header text="Welcome Back!"/>
                    <Subheader text="Enter your Information to login"/>
                    <InputBoxes placeholder="Email" heading="Email"/>
                    <InputBoxes placeholder="Enter your Password" heading="Password"/>
                    <button className="bg-red-500 hover:bg-red-600 text-white font-semibold py-3 rounded-xl transition">
                      Continue
                    </button>
                    <AuthBottom text="Don't have an account?" linktext="Sign Up now" onClick={onClose}/>
                </div>
            </div>
        </>
    )
}

export default LoginPage;