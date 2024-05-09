class CustomError extends Error {
    code: number;
    message: string;

    constructor(code: number, message: string) {
        super(message);
        this.code = code;
        this.message = message;
    }
}

const Errors = {
    Success: new CustomError(0, "Success"),
    ServerError: new CustomError(1, "Server Error"),
    AuthFailed: new CustomError(2, "Auth Failed"),
    ParamsInvalid: new CustomError(3, "Params Invalid"),
    NotFound: new CustomError(4, "Not Found"),
    Exception: new CustomError(100, "Server Exception"),
};

export { Errors, CustomError };;
