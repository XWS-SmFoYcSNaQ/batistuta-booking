using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace user_service.messaging.Configuration
{
    public class NatsConfiguration
    {
        public string HOST { get; set; } = string.Empty;
        public string PORT { get; set; } = string.Empty;
        public string USER { get; set; } = string.Empty;
        public string PASSWORD { get; set; } = string.Empty;
    }
}
