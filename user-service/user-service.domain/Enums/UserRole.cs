using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace user_service.domain.Enums
{
    public enum UserRole
    {
        Nonauthenticated = 0,
        Host = 1,
        Guest = 2
    }
}
