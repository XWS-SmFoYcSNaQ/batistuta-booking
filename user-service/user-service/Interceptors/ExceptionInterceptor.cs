using Grpc.Core;
using Grpc.Core.Interceptors;
using System.Net;

namespace user_service.Interceptors
{
    public class ExceptionInterceptor : Interceptor
    {
        private readonly ILogger<ExceptionInterceptor> _logger;

        public ExceptionInterceptor(ILogger<ExceptionInterceptor> logger)
        {
            _logger = logger;
        }

        public override async Task<TResponse> UnaryServerHandler<TRequest, TResponse>(
            TRequest request,
            ServerCallContext context,
            UnaryServerMethod<TRequest, TResponse> continuation)
        {
            try
            {
                return await continuation(request, context);
            }
            catch (RpcException ex)
            {
                _logger.LogError(ex.ToString());
                throw new RpcException(new Status(ex.StatusCode, ex.Message), ex.Trailers);
            }
        }
    }
}
