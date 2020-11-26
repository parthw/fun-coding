

import org.testng.annotations.Test;


public class MainActivityTests extends BaseAppiumClass{

    @Test(description = "[MainActivity] - Activation Token")
    public void addingTextToActivationToken(){
        androidDriver.findElementByAccessibilityId("Edit_ActivationToken").sendKeys("TIAA-CREF");
        androidDriver.findElementByXPath("/hierarchy/android.widget.FrameLayout/android.widget.LinearLayout/android.widget.FrameLayout/android.widget.LinearLayout/android.widget.FrameLayout/android.widget.FrameLayout/android.view.ViewGroup/android.view.ViewGroup/android.view.ViewGroup/android.view.ViewGroup/android.view.ViewGroup/android.view.ViewGroup[2]/android.view.ViewGroup/android.view.ViewGroup/android.view.ViewGroup/android.view.ViewGroup[1]/android.view.ViewGroup[2]/android.view.ViewGroup").click();
        androidDriver.findElementByAccessibilityId("Activation_Button").click();

        System.out.println("Successfully Executed [MainActivity] - Activation Token...");

    }
}