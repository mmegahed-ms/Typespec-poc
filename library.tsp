import "@typespec/http";

using TypeSpec.Http;

@error
model Error{
    code: int32;
    message: string;
}

interface ResourceInterface<T>{
    @doc ("Get a list of resources")
    @get list(): T[] | Error;
    @get read( @path id: string ): T | Error;
    @put create( @body data: T ): T | Error;
    @patch update(  @body data: T ): T | Error;
    @delete delete( @path id: string ): void | Error;
}