import React from 'react';
import {useQueryParams} from 'hookrouter';

const SearchPage = () => {
    const [queryParams] = useQueryParams();

    const {
        // Use object destructuring and a default value
        // if the param is not yet present in the URL.
        q = ''
    } = queryParams;

    return q
        ? `You searched for "${q}"`
        : 'Please enter a search text';
};

export default SearchPage;
