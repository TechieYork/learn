const increment = () => {
    return {
        type: "INCREMENT",
    };
};

const decrement = () => {
    return {
        type: "DECREMENT",
    };
};

const sign_in = () => {
    return {
        type: "SIGN_IN",
    };
};

const sign_out = () => {
    return {
        type: "SIGN_OUT",
    };
};

export { increment, decrement, sign_in, sign_out };
