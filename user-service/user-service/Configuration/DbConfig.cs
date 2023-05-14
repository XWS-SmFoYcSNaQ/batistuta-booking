using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace user_service.Configuration
{
    public class DbConfig
    {
        public string? Server { get; set; }
        public string? Port { get; set; }
        public string? Database { get; set; }
        public string? User { get; set; }
        public string? Password { get; set; }
    }
}
