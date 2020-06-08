import React from 'react';
import AuthLogin from "../../components/Auth/AuthLogin";
import AuthRegister from "../../components/Auth/AuthRegister";


const HomePage = ({username}) => {
    let isLoggedIn = false
    if (!isLoggedIn) {
        // noinspection HtmlUnknownTarget
        return (
            <div>
                <button onClick={() => <AuthLogin/>}>Login</button>
                <button onClick={() => <AuthRegister/>}>Register</button>
            </div>
        );
    } else {
        return (
            <div>
                This is home page of {username}
            </div>
        );
    }
};

export default HomePage;
