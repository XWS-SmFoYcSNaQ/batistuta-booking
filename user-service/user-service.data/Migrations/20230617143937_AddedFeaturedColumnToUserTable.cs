using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace user_service.data.Migrations
{
    /// <inheritdoc />
    public partial class AddedFeaturedColumnToUserTable : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<bool>(
                name: "Featured",
                table: "Users",
                type: "tinyint(1)",
                nullable: true);

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: new Guid("422bd1e1-6d70-42cc-adc6-89a09b313c01"),
                column: "Featured",
                value: null);

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: new Guid("87abe34c-9935-44c4-aad5-af82f9442c77"),
                column: "Featured",
                value: null);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "Featured",
                table: "Users");
        }
    }
}
