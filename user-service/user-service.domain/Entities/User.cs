﻿using user_service.domain.Enums;

namespace user_service.domain.Entities
{
    public class User
    {
        public Guid Id { get; set; }
        public required UserRole Role { get; set; }
        public required string Username { get; set; }
        public required string Password { get; set; }
        public required string FirstName { get; set; }
        public required string LastName { get; set; }
        public required string Email { get; set; }
        public required string LivingPlace { get; set; }
    }
}
