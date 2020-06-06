import React from 'react';
import {useRoutes} from 'hookrouter';
import HomePage from "./pages/Home/HomePage";
import ProfilePage from "./pages/Profile/ProfilePage";
import Error404 from "./pages/Error/Error404";
import SearchPage from "./pages/Search/SearchPage";
import Tweet from "./components/Tweet/Tweet";


const routes = {
    "/home": () => <HomePage/>,
    "/profile/:username": ({username}) => <ProfilePage username={username}/>,
    "/tweet/:tweetID": ({tweetID}) => <Tweet tweetID={tweetID}/>,
    "/search": () => <SearchPage/>
};

const App = () => {
    const routeResult = useRoutes(routes);

    // noinspection HtmlUnknownTarget
    return (
        <>
            <a href="/home">Home Page</a> <br/>
            <a href="/profile/richard">Profile Page</a> <br/>
            <a href="/tweet/1">Tweet</a> <br/>
            <a href="/search">Search Page</a> <br/>

            {routeResult || <Error404/>}
        </>
    )
}

export default App;
