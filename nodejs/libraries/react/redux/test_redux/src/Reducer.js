const counterReducer = (state = 0, action) => {
    switch (action.type) {
        case "INCREMENT":
            return state + 1;
        case "DECREMENT":
            return state - 1;
        default:
            return state;
    }
};

const loggedReducer = (state = false, action) => {
    switch (action.type) {
        case "SIGN_IN":
            return true;
        case "SIGN_OUT":
            return false;
        default:
            return state;
    }
};

export { counterReducer, loggedReducer };
