package tests;

import org.openqa.selenium.By;
import org.testng.annotations.Test;

public class Tests extends BaseClass{

    @Test
    public void testOne(){
        appiumDriver.findElement(By.id("com.android.chrome:id/search_box_text")).sendKeys("Hacked");
        System.out.println("Completed testOne...");

    }
}
