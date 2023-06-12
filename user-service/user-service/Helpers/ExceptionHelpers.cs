using Grpc.Core;
using MySql.Data.MySqlClient;

namespace user_service.Helpers
{
    public static class ExceptionHelpers
    {
        public static RpcException Handle<T>(this Exception ex,
            ServerCallContext context,
            ILogger<T> logger,
            Guid correlationId) =>
        ex switch
        {
            TimeoutException => HandleTimeoutException((TimeoutException)ex, context, logger, correlationId),
            MySqlException => HandleMySqlException((MySqlException)ex, context, logger, correlationId),
            RpcException => HandleRpcException((RpcException)ex, context, logger, correlationId),
            _ => HandleDefault(ex, context, logger, correlationId)
        };

        private static RpcException HandleTimeoutException<T>(TimeoutException ex,
            ServerCallContext context,
            ILogger<T> logger,
            Guid correlationId)
        {
            logger.LogError(ex, $"CorrelationId: {correlationId} - A timeout occurred");

            var status = new Status(StatusCode.Internal, "An external resource did not answer within the time limit");

            return new RpcException(status, CreateTrailers(correlationId));
        }

        private static RpcException HandleMySqlException<T>(MySqlException ex,
            ServerCallContext context,
            ILogger<T> logger,
            Guid correlationId)
        {
            logger.LogError(ex, $"CorrelationId: ${correlationId} - An SQL error occurred");

            var status = new Status(StatusCode.Internal, "SQL error");

            return new RpcException(status, CreateTrailers(correlationId));
        }


        private static RpcException HandleRpcException<T>(RpcException ex,
            ServerCallContext context,
            ILogger<T> logger,
            Guid correlationId)
        {
            logger.LogError(ex, $"CorrelationId: ${correlationId} - An error occurred");
            var trailers = new Metadata();
            foreach (var trailer in ex.Trailers)
            {
                trailers.Add(trailer);
            }
            trailers.Add("CorrelationId", correlationId.ToString());

            return new RpcException(new Status(ex.Status.StatusCode, ex.Status.Detail), trailers);
        }

        private static RpcException HandleDefault<T>(Exception ex,
            ServerCallContext context,
            ILogger<T> logger,
            Guid correlationId)
        {
            logger.LogError(ex, $"CorrelationId: ${correlationId} - An error occurred");

            return new RpcException(new Status(StatusCode.Internal, ex.Message), CreateTrailers(correlationId));
        }



        private static Metadata CreateTrailers(Guid correlationId)
        {
            var trailers = new Metadata
            {
                { "CorrelationId", correlationId.ToString() }
            };

            return trailers;
        }
    }
}
