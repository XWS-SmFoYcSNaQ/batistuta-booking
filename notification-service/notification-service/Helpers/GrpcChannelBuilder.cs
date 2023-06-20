using Grpc.Net.Client;

namespace notification_service.Helpers
{
    public class GrpcChannelBuilder
    {
        private readonly GrpcChannelOptions _grpcChannelOptions;

        public GrpcChannelBuilder(GrpcChannelOptions grpcChannelOptions)
        {
            _grpcChannelOptions = grpcChannelOptions;
        }

        public GrpcChannel Build(string url)
        {
            var address = url.StartsWith("http") ? url : $"dns:///{url}";
            return GrpcChannel.ForAddress(new Uri(address), _grpcChannelOptions);
        }
    }
}
