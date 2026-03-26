import { useState, type JSX } from "react";
import Sidebar from "../components/SideBar";
import Appbar from "../components/Appbar";
import Header from "../components/Header";
import Subheader from "../components/Subheading";
import LoginPage from "./Login";
import SignupPage from "./Signup";

// global states
import { useAuthStore } from "../store/AuthStat";
import { UseCatagoryProp } from "../store/PinProps";

const Ideas = ():JSX.Element=>{
    const [loginOpen,setLoginOpen]=useState(false);
    const [signupOpen,setSignupOpen]=useState(false);
    const LoggedIn = useAuthStore((s)=>s.isLoggedIn);

    const [Catagories]=useState(['Travel', 'Photography', 'Food', 'Design', 'Music', 
    'Art', 'Technology', 'Sports', 'Fashion', 'Gaming']);
    const [activeCatagory,setActiveCatagory]=useState("");

    return(<>
    <div className="flex h-screen">
        {LoggedIn && <LoginPage onClose={()=>{}} openSignup={()=>{ // add ! when doing state management for test purposes removed
            setLoginOpen(false);
            setSignupOpen(true);
        }}/>}
        {loginOpen && <LoginPage onClose={()=>setLoginOpen(false)} openSignup={()=>{
            setLoginOpen(false);
            setSignupOpen(true);
        }}/>}
        {signupOpen && <SignupPage onClose={()=>{
            setSignupOpen(false);
        }} openLogin={()=>{
            setSignupOpen(false);
            setLoginOpen(true);
        }}/>}
        
            <Sidebar/>
        <div className=" relative flex flex-col items-center w-full gap-5">
            <Appbar/>
            <Header text="Explore Your Ideas!"/>
            <Subheader text="Search the images based on your interests and hobbies :) "/>

            {/* Catagories component */}
            <div className="w-full">
                <div className="overflow-x-auto scrollbar-hide">
                    <div className="flex gap-3 px-6 pb-4 min-w-min">
                        {Catagories.map((category)=>(
                            <button key={category}
                            onClick={()=>setActiveCatagory(category)}
                            className={`px-6 py-2 rounded-full whitespace-nowrap font-medium transition-all ${
                            activeCatagory === category
                            ? 'bg-red-500 text-white'
                            : 'bg-gray-200 text-gray-800 hover:bg-gray-300'
                            }`}
                            >
                            {category}
                            </button>
                        ))}
                    </div>
                </div>   
            </div>

            {/* Render Pins */}
        </div>
    </div>
    </>)
}

export default Ideas;