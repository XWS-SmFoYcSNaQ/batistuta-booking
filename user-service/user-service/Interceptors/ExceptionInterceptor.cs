using Grpc.Core;
using Grpc.Core.Interceptors;
using System.Net;
using user_service.Helpers;

namespace user_service.Interceptors
{
    public class ExceptionInterceptor : Interceptor
    {
        private readonly ILogger<ExceptionInterceptor> _logger;
        private readonly Guid _correlationId;

        public ExceptionInterceptor(ILogger<ExceptionInterceptor> logger)
        {
            _logger = logger;
            _correlationId = Guid.NewGuid();
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
                throw ex.Handle(context, _logger, _correlationId);
            }
        }
    }
}
